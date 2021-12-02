# Test Driven Development & GoLang Practice

Exercise I performed to practice GoLang and Test Driven Development based on these requirements: <https://github.com/testdouble/contributing-tests/wiki/Greeting-Kata>

This Kata is designed to help practice how a test of a pure function ought to look. It is intentionally designed to start with a very easy, non-branching base case which slowly becomes addled with complexity as additional requirements are added that will require significant branching and eventually a pressure to compose additional units.

I created a minimal wrapper around GoLang's built in testing framework so I could supply a simple list of TestCase objects that were easy to construct with the following parameters:

- Name: the name of the test so I can keep track of what I'm actually testing
- Inputs: parameters to the function I am testing
- Exp: expected output
- Pass: if the test is "supposed" to pass or not. Not currently used, but would be useful for ensuring that unwanted outputs don't make their way past tests.
