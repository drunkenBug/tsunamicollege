package database

var(
)

type RequestData struct {
        Op string `json:"op"`
        ID uint  `json:"id"`
}

type InsertData struct {
        Op string `json:"op"`
        Title string `json:"title"`
        Solution string `json:"solution"`
}

type Question struct {
        ID uint `json:"id"`
        Title    string `json:"title"`
        Solution string `json:"solution"`
}


