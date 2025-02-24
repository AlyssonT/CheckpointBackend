using Checkpoint.API.Responses;
using Microsoft.AspNetCore.Mvc;

namespace Checkpoint.API.Controllers
{
    [Route("[controller]")]
    [ApiController]
    public class HelloController : ControllerBase
    {
        [HttpGet]
        public IActionResult Hello() {
            var response = ResponseDto.CreateSuccess("Hello World!");
            return Ok(response);
        }
    }
}
