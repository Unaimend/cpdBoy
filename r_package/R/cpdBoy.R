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
  print(typeof(data))
  if (typeof(data) == "character") {
    # Create a list representing the JSON body
    json_body <- list(
      text = data
    )
  }

  if (is.vector(data)) {
    # Create a list representing the JSON body
    str <- paste(data, collapse = ",")
    json_body <- list(
      text = str
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
  
  if (is.vector(data)) {
    df <- read.csv(textConnection(content(response, "text")), header = FALSE)
    colnames(df) <- c("id", "name")
  }
  # Print the response
  return(df)
}
