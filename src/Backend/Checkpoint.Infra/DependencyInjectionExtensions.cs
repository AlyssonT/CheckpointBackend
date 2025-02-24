
using Checkpoint.Domain.Repositories;
using Checkpoint.Infra.Repositories;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.DependencyInjection;

namespace Checkpoint.Infra;

public static class DependencyInjectionExtensions
{
    public static void AddInfra(this IServiceCollection services)
    {
        services.AddScoped<IUserRepository, UserRepository>();
        services.AddDbContext<CheckpointDbContext>();
    }
}
