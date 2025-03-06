using Checkpoint.Application.Handlers.User;
using Checkpoint.Application.Services;
using Checkpoint.Exceptions;
using Checkpoint.Exceptions.ExceptionBase;
using CommonTestUtilities.Repositories;
using CommonTestUtilities.Requests;

namespace Handlers.Test.User;

public class CreateUserHandlerTest
{
    [Fact]
    public async Task Success()
    {
        var command = CreateUserCommandBuilder.Build();
        var handler = CreateHandler();

        var result = await handler.Handle(command, CancellationToken.None);
        Assert.NotNull(result);
        Assert.True(result.Id >= 0);
    }

    [Fact]
    public async Task NameOrEmailAlreadyExists()
    {
        var command = CreateUserCommandBuilder.Build();
        var handler = CreateHandler(command.Email, command.Name);

        Task<CreatedUserData> func() => handler.Handle(command, CancellationToken.None);
        var result = await Assert.ThrowsAsync<UserAlreadyExistsException>(func);

        Assert.NotNull(result);
        Assert.Equal(MessagesExceptions.EMAIL_OR_NAME_ALREADY_REGISTERED, result.Message);
    }

    private static CreateUserHandler CreateHandler(string? email = null, string? name = null)
    {
        var userRepositoryBuilder = new UserRepositoryBuilder();

        if (email is not null && name is not null)
        {
            userRepositoryBuilder.UserExists(email, name);
        }

        return new CreateUserHandler(userRepositoryBuilder.Build(), new PasswordEncrypter("ABC"));
    }
}
