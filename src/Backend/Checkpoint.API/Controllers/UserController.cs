using Checkpoint.API.Responses;
using Checkpoint.Application.Commands.User;
using MediatR;
using Microsoft.AspNetCore.Mvc;

namespace Checkpoint.API.Controllers
{
    [Route("[controller]")]
    [ApiController]
    public class UserController : ControllerBase
    {
        private readonly IMediator _mediator;

        public UserController(IMediator mediator)
        {
            _mediator = mediator;
        }

        [HttpPost]
        public async Task<IActionResult> Register([FromBody] CreateUserCommand command)
        {
            var result = await _mediator.Send(command);
            var response = ResponseDto.CreateSuccess(result, StatusCodes.Status201Created);
            return Created("", response);
        }
    }
}
