
package models

type User struct {
    ID       int64   `json:"id"`
    Username string `json:"username" binding:"required,min=3,max=30,alphanum"`
    Email    string `json:"email" binding:"required,email,max=254"`
    Password string `json:"password" binding:"required,min=8,max=72"`
}
