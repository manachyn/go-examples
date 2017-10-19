server provides the main function and the handler for /search.
userip provides functions for extracting a user IP address from a request and associating it with a Context.
google provides the Search function for sending a query to Google.

The server program handles requests like /search?q=golang by serving the first few Google search results for golang. It registers handleSearch to handle the /search endpoint.


go run server.go