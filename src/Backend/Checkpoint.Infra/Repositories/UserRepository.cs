using Checkpoint.Domain.Entities;
using Checkpoint.Domain.Repositories;
using Microsoft.EntityFrameworkCore;

namespace Checkpoint.Infra.Repositories;

public class UserRepository : IUserRepository
{
    private readonly CheckpointDbContext _context;
    public UserRepository(CheckpointDbContext context)
    {
        _context = context;
    }
    public async Task<long> CreateUser(string name, string password, string email)
    {
        var result = await _context.Users.AddAsync(new User
        {
            Name = name,
            Password = password,
            Email = email
        });

        await _context.SaveChangesAsync();
        return result.Entity.Id;
    }

    public async Task<User> GetUserById(long id)
    {
        var user = await _context.Users.FindAsync(id);

        if (user is null)
        {
            var listUsers = await _context.Users.ToListAsync();
            foreach (var item in listUsers)
            {
                Console.WriteLine(item);
            }
        }

        return user;
    }
}
