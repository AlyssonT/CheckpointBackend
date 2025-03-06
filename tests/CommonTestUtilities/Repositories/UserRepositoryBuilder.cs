using Checkpoint.Domain.Repositories;
using Moq;

namespace CommonTestUtilities.Repositories;

public class UserRepositoryBuilder
{
    private readonly Mock<IUserRepository> _userRepository;

    public UserRepositoryBuilder()
    {
        _userRepository = new Mock<IUserRepository>();
    }

    public void UserExists(string email, string name)
    {
        _userRepository.Setup(repository => repository.UserExists(email, name)).ReturnsAsync(true);
    }

    public IUserRepository Build()
    {
        return _userRepository.Object;
    }
}
