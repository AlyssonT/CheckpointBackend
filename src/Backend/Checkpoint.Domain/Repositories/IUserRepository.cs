using Checkpoint.Domain.Entities;

namespace Checkpoint.Domain.Repositories;

public interface IUserRepository
{
    Task<long> CreateUser(string name, string password, string email);
    Task<User> GetUserById(long id);
}
