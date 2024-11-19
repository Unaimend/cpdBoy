library(httr)
library(jsonlite)
# Load the 'httr' library

#' @param data Either a string, list or data.frame.as.data.frame.
#' In the case of a data frame it should contain `cpd` either in its colnames or the first row, 
#' if thats where your headers are.
#' @export
cpd <- function(data) {

  # Define the API endpoint
  url <- "127.0.0.1:3000/message"

  if (typeof(data) == "character") {
    # Create a list representing the JSON body
    json_body <- list(
      text = data
    )
  }

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
  return(content(response, "text"))
}
