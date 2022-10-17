# Cloud Humans - Backend Take Home Assignment

People are the fuel that moves Cloud Humans. In order to speed up the company we need to be great at pairing people (we call them Pros) to problems as fast as possible. 

Cloud Humans determines which project a new Pro will tackle based on the answers they provide in the Pro Portal (a frontend application). The frontend send this Pro's **application** to an API that runs an algorithm upon the data and returns the project in which the Pro will be allocated. 

For this assignment, you will create a simple version of that application by coding an API endpoint that receives a JSON payload with the Pro application and returns a JSON response with the selected project – you don’t have to worry about the frontend of the application. 

## The input
The would-be frontend of this application would ask the Pros for their personal information and then it would send to your API a payload following this contract:

```JSON
{
  "age": 35,
  "education_level": "high_school",
  "past_experiences": {
    "sales": false,
    "support": true
  },
  "internet_test": {
    "download_speed": 50.4,
    "upload_speed": 40.2
  },
  "writing_score": 0.6,
  "referral_code": "token1234"
}
```

### Pro Application
All attributes are required except for the referral_code.

- Age: an integer equal or greater than 0.
- Education Level: `"no_education"`, `"high_school"` or `"bachelors_degree_or_high"`)
- Past experiences: JSON object with bool attributes `sales` and `support`
- Internet Test: JSON object with the result of an internet test already made by the frontend (the attributes are positive floats, where 50.4 == 50.4 megabytes)
- Writing Score: float between 0 and 1
- Referral code: a string with a referral token provided by another Pro

## The algorithm
The application receives the JSON payload through the API endpoint and returns the **project** which the **Pro** will be paired with, based on the algorithm detailed below.

First, it calculates an **eligibility score** by giving points to each information provided, based on these current rules:


1. If the Pro is under age, she is ineligible to be paired with any project
2. If the Pro education level is "high_school" she receives 1 score point, if it is "bachelors_degree_or_high" she receives 2
3. If the Pro has experience with sales and/or support she receives 5 and 3 score points respectively
4. If her internet download speed is greater than 50mb she receives 1 score point, if it is lower than 5mb her score points are deducted in 1. The same rule applies for upload speed. 
5. If her writing score is lower than 0.3 she is deducted in 1 point. If between 0.3 and 0.7 she receives 1 point. If greather than 0.7 she receives 2 points.
6. If the referral code is valid (consider the value "token1234" as valid) she receives 1 score point.

With the **final score** resulting from this algorithm we have to determine which projects the Pro is eligible for, as each project has it's own complexity and requires different skills. And then we pair the Pro with the first critical project she is eligible for, based on this sorted **project list** below: 

1. The project `"Calculate the Dark Matter of the universe for Nasa"` requires Pros with score greater than 10
2. The project `"Determine if the Schrodinger's cat is alive"` requires Pros with score greater than 5 
3. The project `"Attend to users support for a YXZ Company"` requires Pros with score greater than 3 
4. The project `"Collect specific people information from their social media for XPTO Company"` requires Pros with score greater than 2

## The output
Considering the data provided above, the application should return the following JSON payload:

```JSON
{
    "score": 7,
    "selected_project": "determine_schrodinger_cat_is_alive",
    "eligible_projects": ["determine_schrodinger_cat_is_alive", "support_users_from_xyz", "collect_information_for_xpto"],
    "ineligible_projects": ["calculate_dark_matter_nasa"]
}
```

## Criteria
You may use any language and framework provided that you build a solid system with an emphasis on code quality, simplicity, readability, maintainability, and reliability, particularly regarding architecture and testing.

Be aware that we will mainly take into consideration the following evaluation criteria:
* How clean and organized your code is;
* If you implemented the business rules correctly;
* How good your automated tests are (qualitative over quantitative).

Other important notes:
* Develop an extensible algorithm
* Add to the README file: (1) instructions to run the code; (2) what were the main technical decisions you made; (3) relevant comments about your project 
* You must use English in your code and also in your docs

You can take the time you need for implementing this assignment, but we ask you to stop the commits at least 1 day before our interview so we can review it. We expect you to learn fast, communicate with us if you need, and make decisions regarding its implementation & scope to achieve the expected results on time.

