using Checkpoint.Application.Commands.User;
using Checkpoint.Application.Services;
using Checkpoint.Domain.Repositories;
using Checkpoint.Exceptions;
using Checkpoint.Exceptions.ExceptionBase;
using FluentValidation;
using MediatR;

namespace Checkpoint.Application.Handlers.User;

public class CreateUserHandler : IRequestHandler<CreateUserCommand, CreatedUserData>
{
    private readonly PasswordEncrypter _passwordEncrypter;
    private readonly IUserRepository _userRepository;
    public CreateUserHandler(IUserRepository userRepository, PasswordEncrypter passwordEncrypter)
    {
        _passwordEncrypter = passwordEncrypter;
        _userRepository = userRepository;
    }
    public async Task<CreatedUserData> Handle(CreateUserCommand request, CancellationToken cancellationToken)
    {
        await Validate(request);

        var passwordHash = _passwordEncrypter.Encrypt(request.Password);
        var userId = await _userRepository.CreateUser(request.Name, passwordHash, request.Email);

        return new CreatedUserData { Id = userId };
    }

    private async Task Validate(CreateUserCommand request)
    {
        var validator = new CreateUserValidator();
        var result = await validator.ValidateAsync(request);
        if (!result.IsValid)
        {
            var errorMessages = result.Errors.Select(x => x.ErrorMessage).ToList();
            throw new ErrorOnValidationException(errorMessages);
        }

        if (await _userRepository.UserExists(request.Email, request.Name))
        {
            throw new UserAlreadyExistsException(MessagesExceptions.EMAIL_OR_NAME_ALREADY_REGISTERED);
        }
    }
}

public class CreatedUserData()
{
    public long Id { get; set; }
}

public class CreateUserValidator : AbstractValidator<CreateUserCommand>
{
    public CreateUserValidator()
    {
        RuleFor(x => x.Name).NotEmpty().WithMessage(MessagesExceptions.NAME_EMPTY);
        RuleFor(x => x.Password.Length).GreaterThanOrEqualTo(6).WithMessage(MessagesExceptions.PASSWORD_INVALID);
        RuleFor(x => x.Email).EmailAddress().WithMessage(MessagesExceptions.EMAIL_INVALID);
    }
}
