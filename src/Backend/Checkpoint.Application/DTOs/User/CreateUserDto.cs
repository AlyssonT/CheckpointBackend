namespace Checkpoint.Application.DTOs.User;

public class CreateUserDto
{
    public required string Name { get; set; }
    public required string Password { get; set; }
    public required string Email { get; set; }
}
