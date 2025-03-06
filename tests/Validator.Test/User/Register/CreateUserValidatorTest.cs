using Checkpoint.Application.Handlers.User;
using Checkpoint.Exceptions;
using CommonTestUtilities.Requests;

namespace Validator.Test.User.Register;

public class CreateUserValidatorTest
{
    [Fact]
    public void Success()
    {
        var validator = new CreateUserValidator();
        var command = CreateUserCommandBuilder.Build();

        var result = validator.Validate(command);

        Assert.True(result.IsValid);
    }

    [Fact]
    public void NameEmpty()
    {
        var validator = new CreateUserValidator();
        var command = CreateUserCommandBuilder.Build();
        command.Name = "";

        var result = validator.Validate(command);

        Assert.False(result.IsValid);
        Assert.Single(result.Errors);
        Assert.Equal(MessagesExceptions.NAME_EMPTY, result.Errors[0].ErrorMessage);
    }

    [Theory]
    [InlineData("email")]
    [InlineData("email@email")]
    [InlineData("@email.com")]
    public void EmailInvalid(string email)
    {
        var validator = new CreateUserValidator();
        var command = CreateUserCommandBuilder.Build();
        command.Email = email;

        var result = validator.Validate(command);

        Assert.False(result.IsValid);
        Assert.Single(result.Errors);
        Assert.Equal(MessagesExceptions.EMAIL_INVALID, result.Errors[0].ErrorMessage);
    }

    [Fact]
    public void PasswordInvalid()
    {
        var validator = new CreateUserValidator();
        var command = CreateUserCommandBuilder.Build();
        command.Password = "123ab";

        var result = validator.Validate(command);

        Assert.False(result.IsValid);
        Assert.Single(result.Errors);
        Assert.Equal(MessagesExceptions.PASSWORD_INVALID, result.Errors[0].ErrorMessage);
    }
}
