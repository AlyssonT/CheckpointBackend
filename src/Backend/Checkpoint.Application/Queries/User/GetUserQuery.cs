using Checkpoint.Application.DTOs.User;
using MediatR;

namespace Checkpoint.Application.Queries.User;

public class GetUserQuery : IRequest<UserDto>
{
    public long Id { get; set; }
    public GetUserQuery(long id)
    {
        Id = id;
    }
}
