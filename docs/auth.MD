# Authentication API Dokumentasi
Authentication ini digunakan untuk meminta akess ke server berdasarkan kredential pengguna  
Endpont:
- ### Login
    - URL: `/v1/auth`
    - Method: POST
    - Body:
        ```json
       {
          "username": "string",
          "password": "string"
       }
        ```
    - Response:
        ```json
        {
          "status": "success",
          "data": {
            "access_token": "string",
            "refresh_token": "string"
          }
        }
        ```
- ### Refresh Token
    - URL: `/v1/auth`
    - Method: PUT
    - Body:
        ```json
       {
          "refresh_token": "string"
       }
        ```
    - Response:
        ```json
        {
          "status": "success",
          "data": {
            "access_token": "string"
          }
        }
        ```
- ### Logout
    - URL: `/v1/auth`
    - Method: DELETE
    - Body:
        ```json
       {
          "refresh_token": "string"
       }
        ```
    - Response:
        ```json
        {
          "status": "success",
          "message": "string"
        }
        ```
