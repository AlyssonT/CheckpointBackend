using System.Net;
using System.Net.Http.Json;
using System.Text.Json;
using Checkpoint.Application.Commands.User;
using CommonTestUtilities.Requests;

namespace WebApi.Test.User;

public class CreateUserTest : IClassFixture<CustomWebApplicationFactory>
{
    private readonly HttpClient _httpClient;
    public CreateUserTest(CustomWebApplicationFactory factory)
    {
        _httpClient = factory.CreateClient();
    }

    [Fact]
    public async Task Success()
    {
        var command = CreateUserCommandBuilder.Build();
        var response = await _httpClient.PostAsJsonAsync("User", command);

        Assert.Equal(HttpStatusCode.Created, response.StatusCode);

        await using var responseBody = await response.Content.ReadAsStreamAsync();
        var responseData = await JsonDocument.ParseAsync(responseBody);

        var success = responseData.RootElement.GetProperty("success").GetBoolean();
        Assert.True(success);

        var statusCode = responseData.RootElement.GetProperty("statusCode").GetInt32();
        Assert.Equal((int)HttpStatusCode.Created, statusCode);
    }

    [Fact]
    public async Task InvalidFields()
    {
        var command = new CreateUserCommand
        {
            Name = "",
            Email = "invalidEmail",
            Password = "123",
        };

        var response = await _httpClient.PostAsJsonAsync("User", command);

        Assert.Equal(HttpStatusCode.BadRequest, response.StatusCode);

        await using var responseBody = await response.Content.ReadAsStreamAsync();
        var responseData = await JsonDocument.ParseAsync(responseBody);

        var success = responseData.RootElement.GetProperty("success").GetBoolean();
        Assert.False(success);

        var statusCode = responseData.RootElement.GetProperty("statusCode").GetInt32();
        Assert.Equal((int)HttpStatusCode.BadRequest, statusCode);

        var messagesLength = responseData.RootElement.GetProperty("messages").GetArrayLength();
        Assert.Equal(3, messagesLength);
    }

    [Fact]
    public async Task EmailOrNameAlreadyExists()
    {
        var command = CreateUserCommandBuilder.Build();

        var response = await _httpClient.PostAsJsonAsync("User", command);
        Assert.Equal(HttpStatusCode.Created, response.StatusCode);

        var originalEmail = command.Email;

        command.Email = "another@email.com";
        response = await _httpClient.PostAsJsonAsync("User", command);
        Assert.Equal(HttpStatusCode.Conflict, response.StatusCode);

        command.Email = originalEmail;
        command.Name = "anotherName";
        response = await _httpClient.PostAsJsonAsync("User", command);
        Assert.Equal(HttpStatusCode.Conflict, response.StatusCode);
    }
}
