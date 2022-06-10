package domain

import (
    "strconv"
    "strings"
)

type Post struct {
    Id     int32  `json:"id,omitempty"`
    UserId int32  `json:"userId,omitempty"`
    Title  string `json:"title,omitempty"`
    Body   string `json:"body,omitempty"`
}

func (p *Post) String() string {
    var str strings.Builder
    str.WriteString("Post {")
    str.WriteString("id: " + strconv.Itoa(int(p.Id)))
    str.WriteString("; user_id: " + strconv.Itoa(int(p.UserId)))
    str.WriteString("; title: " + shorten(p.Title))
    str.WriteString("; body: " + shorten(p.Body) + "}")

    return str.String()
}

func shorten(s string) string {
    if len(s) > 20 {
        return s[0:20] + "..."
    }

    return s
}
