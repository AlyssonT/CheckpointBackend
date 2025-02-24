using System;

namespace Checkpoint.Domain.Entities;

public class User
{
    public long Id { get; set; }
    public string Name { get; set; } = "";
    public string Email { get; set; } = "";
    public string Password { get; set; } = "";
    public DateTime CreatedAt { get; set; } = DateTime.Now;
    public bool Active { get; set; } = true;
}
