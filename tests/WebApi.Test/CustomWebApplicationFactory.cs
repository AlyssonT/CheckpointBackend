using Checkpoint.Infra;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Mvc.Testing;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.DependencyInjection;

namespace WebApi.Test;

public class CustomWebApplicationFactory : WebApplicationFactory<Program>
{
    protected override void ConfigureWebHost(IWebHostBuilder builder)
    {
        builder.UseEnvironment("Test")
            .ConfigureServices(services =>
            {
                var descriptor = services.SingleOrDefault(d => d.ServiceType == typeof(DbContextOptions<CheckpointDbContext>));
                if (descriptor is not null)
                {
                    services.Remove(descriptor);
                }

                var provider = services.AddEntityFrameworkInMemoryDatabase().BuildServiceProvider();
                services.AddDbContext<CheckpointDbContext>(options => {
                   options.UseInMemoryDatabase("InMemoryDbForTesting");
                   options.UseInternalServiceProvider(provider);
                });
            });
    }
}
