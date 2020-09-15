# Welcome to the word pyramid webserver

## Setup 
1. Clone the repo
2. Using the CLI, Run `make run` from the base `webserver-pyramid` directory 
3. Using Postman or cURL make a request to the local server. API inforamation below

## API
```
Path: /pyramid
Method: Post
Params: 
 - show_details 
    type: <bool> 
    location: query params
    function: Shows the detailed version of the response
- Words
    type: array
    location: in body
    fuction: The words to be checked
    Example:
        {
            "words": [
                "banana",
                "boo",
                "b"
            ]
        }
Response: 
    type: json
    example (without Details): 
    {
        "results": [
            {
                "is_pyramid": false,
                "word": "bs",
                "ShowDetails": false
            },
            {
                "is_pyramid": true,
                "word": "banana",
                "ShowDetails": false
            }
        ]
    }
    example (with Details): 
    {
        "results": [
            {
                "characters": [
                    {
                        "letter": "b",
                        "number_of_times_seen": 1
                    },
                    {
                        "letter": "s",
                        "number_of_times_seen": 1
                    }
                ],
                "is_pyramid": false,
                "word": "bs",
                "ShowDetails": false
            },
            {
                "characters": [
                    {
                        "letter": "b",
                        "number_of_times_seen": 1
                    },
                    {
                        "letter": "n",
                        "number_of_times_seen": 2
                    },
                    {
                        "letter": "a",
                        "number_of_times_seen": 3
                    }
                ],
                "is_pyramid": true,
                "word": "banana",
                "ShowDetails": false
            }
        ]
    }
```

## Assumptions
1. In order for a word to be considered a word pyramid, the word can not be more than 3 characters
2. Special Characters and numbers can be apart of a "word"
3. Case Insentive. "S" == "s" 


*Disclaimer: Please excuse the typos. This was a fun project to work on. It's doint much more than the required but it's tough to stop when the creative juices start flowing. 