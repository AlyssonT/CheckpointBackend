using Bogus;
using Checkpoint.Application.Commands.User;

namespace CommonTestUtilities.Requests;

public static class CreateUserCommandBuilder
{
    public static CreateUserCommand Build()
    {
        return new Faker<CreateUserCommand>()
            .RuleFor(user => user.Name, f => f.Person.FirstName)
            .RuleFor(user => user.Email, (f, u) => f.Internet.Email(u.Name))
            .RuleFor(user => user.Password, f => f.Internet.Password());
    }
}
