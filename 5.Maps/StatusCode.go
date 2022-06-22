package statuscode

const (
	Deleted  = 00
	NotFound = 02
)

var statusText = map[int]string{

	Deleted:  "Data Succesfully Deleted",
	NotFound: "Data Not Found",
}

func StatusText(code int) string {
	return statusText[code]
}
