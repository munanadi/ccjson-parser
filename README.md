## ccjsonparser

A json parser written in golang. From this challenge [here](https://codingchallenges.fyi/challenges/challenge-json-parser/)

---

1. Using a stack to match brackets
2. Pushing Open brakcets and popping off closed ones.
3. Mismatch is invalid JSON else valid JSON.

---

1. So we are given a json input
2. We need to iterate over that input and create a slice of tokens
3. We operate a lexer over this slice of tokens
4. We can know if the passed JSON is valid or not

---

### Usage

```sh
make run
```

---

### Reading

1. [Building a JSON Parser and Query Tool with Go](https://medium.com/@bradford_hamilton/building-a-json-parser-and-query-tool-with-go-8790beee239a)