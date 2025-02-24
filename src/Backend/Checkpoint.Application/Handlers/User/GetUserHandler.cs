using Checkpoint.Application.DTOs.User;
using Checkpoint.Application.Queries.User;
using Checkpoint.Domain.Repositories;
using Checkpoint.Exceptions.ExceptionBase;
using MediatR;

namespace Checkpoint.Application.Handlers.User;

public class GetUserHandler : IRequestHandler<GetUserQuery, UserDto>
{
    private readonly IUserRepository _userRepository;
    public GetUserHandler(IUserRepository userRepository)
    {
        _userRepository = userRepository;
    }
    public async Task<UserDto> Handle(GetUserQuery request, CancellationToken cancellationToken)
    {
        var user = await _userRepository.GetUserById(request.Id);

        if (user is null)
        {
            throw new NotFoundException("User not found");
        }

        return new UserDto
        {
            Id = user.Id,
            Name = user.Name,
            Email = user.Email,
            Password = user.Password,
        };
    }
}
