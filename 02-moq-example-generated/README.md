# Generated Example

With the same service and interface as 01 we now auto generate the mocked struct using [Moq](https://github.com/matryer/moq). This is a code generator that takes an interface and generates mocked implementations which we can use to define behaviour in our tests. It also tracks calls to each method, so you can ensure the mock was called the correct amount of times.

To install, follow the instructions at the above link.

This outputs a bit more code than our simple mocked struct we made ourselves, it is also thread safe when logging the calls using the sync package to lock the slice.

The main advantage of this is it takes this job of you manually generating mocks away from you and places it at the responsibility of the machine. Mock generation is boring and you will need to go back to all your mocked objects and amend them as and when your interface contracts change or get added to. This is tedious.

To run the mock execute in the cli `moq -out airplane_mock.go . AirplaneReader`

The next example 03, shows how to make this task even easier with GoGenerate.
