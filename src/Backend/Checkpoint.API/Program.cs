using System.Reflection;
using System.Text.Json.Serialization;
using Checkpoint.API.Filters;
using Checkpoint.Application;
using Checkpoint.Infra;
using Scalar.AspNetCore;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddOpenApi();
builder.Services.AddControllers()
    .AddJsonOptions(options => { options.JsonSerializerOptions.Converters.Add(new JsonStringEnumConverter()); });

builder.Services.AddMvc(options => options.Filters.Add<ExceptionFilter>());
builder.Services.AddMediatR(cfg => cfg.RegisterServicesFromAssembly(Assembly.Load("Checkpoint.Application")));
builder.Services.AddApplication(builder.Configuration);
builder.Services.AddInfra();

var app = builder.Build();

if (app.Environment.IsDevelopment())
{
    app.MapOpenApi();
    app.MapScalarApiReference(options =>
    {
        options
            .WithTitle("Checkpoint API")
            .WithTheme(ScalarTheme.Mars)
            .WithDefaultHttpClient(ScalarTarget.CSharp, ScalarClient.HttpClient);
    });
}

app.UseCors("AllowAnyOrigin");
app.MapControllers();

app.Run();
