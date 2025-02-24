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

    public async Task<bool> UserExists(string email, string name)
    {
        return await _context.Users.AnyAsync(x => x.Email == email || x.Name == name);
    }
}
