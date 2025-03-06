using Checkpoint.Domain.Entities;
using Microsoft.EntityFrameworkCore;

namespace Checkpoint.Infra;

public class CheckpointDbContext(DbContextOptions<CheckpointDbContext> options) : DbContext(options)
{
    public DbSet<User> Users { get; set; }
    public string DbPath { get; } = Path.Join(Environment.GetFolderPath(Environment.SpecialFolder.LocalApplicationData), "Checkpoint.db");

    protected override void OnConfiguring(DbContextOptionsBuilder options)
    {
        options.UseSqlite($"Data Source={DbPath}");
    }
}
