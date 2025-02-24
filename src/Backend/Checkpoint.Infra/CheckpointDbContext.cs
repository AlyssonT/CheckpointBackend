using Checkpoint.Domain.Entities;
using Microsoft.EntityFrameworkCore;

namespace Checkpoint.Infra;

public class CheckpointDbContext : DbContext
{
    public DbSet<User> Users { get; set; }
    public string DbPath { get; }
    public CheckpointDbContext()
    {
        var folder = Environment.SpecialFolder.LocalApplicationData;
        var path = Environment.GetFolderPath(folder);
        DbPath = Path.Join(path, "Checkpoint.db");
    }
    protected override void OnConfiguring(DbContextOptionsBuilder options)
    {
        options.UseSqlite($"Data Source={DbPath}");
    }
}
