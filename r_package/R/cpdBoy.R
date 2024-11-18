# Load the 'httr' library
#' Does things
#' @export
d <- function() {
  library(httr)
  library(jsonlite)
  
  # Define the API endpoint
  url <- "127.0.0.1:3000/message"
  
  # Create a list representing the JSON body
  json_body <- list(
    text = "cpd00058"
  )
  
  # Convert the list to JSON
  json_body <- toJSON(json_body, auto_unbox = TRUE)
  
  # Send the POST request
  response <- POST(
    url, 
    body = json_body, 
    encode = "json", 
    add_headers("Content-Type" = "application/json")
  )
  
  # Print the response
  print(content(response, "text"))
}