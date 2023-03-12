package weibo

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	blank      = "$nbsp;"
	colon      = ":"
	whitespace = ""
)

func toStringReplace(text string) string {
	return strings.ReplaceAll(strings.TrimPrefix(string(text), colon), blank, whitespace)
}

func toStringReg(text string) string {
	re := regexp.MustCompile(`&nbsp;`)
	return re.ReplaceAllStringFunc(text, func(s string) string {
		return strings.Repeat(" ", len(s))
	})
}

func toUrlValue(text string) string {
	u, _ := url.Parse(text)
	path := strings.Split(u.Path, "/")
	return path[len(path)-1]
}

func toUrlUid(text string) string {
	u, _ := url.Parse(text)
	value, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return ""
	}
	return value.Get("uid")
}

func toDate(text string) time.Time {
	text = strings.TrimSpace(text)
	now := time.Now()
	switch {
	case text == "刚刚":
		return now
	case text[len(text)-3:] == "分钟前":
		minutesAgo := text[:len(text)-3]
		minutes, _ := strconv.Atoi(minutesAgo)
		return now.Add(time.Duration(-minutes) * time.Minute)
	case text[len(text)-3:] == "小时前":
		hoursAgo := text[:len(text)-3]
		hours, _ := strconv.Atoi(hoursAgo)
		return now.Add(time.Duration(-hours) * time.Hour)
	case len(text) == len("2022-09-12 17:04:49"):
		v, _ := time.Parse("2006-01-02 15:04:05", text)
		return v
	case strings.Contains(text, "今天"):
		hours := strings.TrimSpace(strings.Replace(text, "今天", "", -1))
		l := strings.Split(hours, ":")
		hour := l[0]
		min := l[1]
		hourr, _ := strconv.Atoi(hour)
		minr, _ := strconv.Atoi(min)
		y, m, d := now.Date()
		return time.Date(y, m, d, hourr, minr, 0, 0, time.Local)
	default:
		var t time.Time
		if strings.Contains(text, "年") {
			t, _ = time.Parse("2006年01月02日 15:04", text)
		} else if strings.Contains(text, "-") {
			t, _ = time.Parse("2006-01-02 15:04:05", text)
		} else {
			t, _ = time.Parse("01月02日 15:04", text)
			t = t.AddDate(now.Year(), 0, 0)
		}
		return t
	}
}
