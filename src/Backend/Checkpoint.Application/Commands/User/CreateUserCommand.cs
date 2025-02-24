using System;
using Checkpoint.Application.Handlers.User;
using MediatR;

namespace Checkpoint.Application.Commands.User;

public class CreateUserCommand : IRequest<CreatedUserData>
{
    public required string Name { get; set; }
    public required string Password { get; set; }
    public required string Email { get; set; }
}
