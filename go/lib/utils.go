package lib

type Utils struct {}

func NewUtils() Utils {
    return Utils{}
}

func (this Utils) Max(v1 int, v2 int) int {
    if v1 > v2 {
        return v1
    } else {
        return v2
    }
}

func (this Utils) Min(v1 int, v2 int) int {
    if v1 < v2 {
        return v1
    } else {
        return v2
    }
}
