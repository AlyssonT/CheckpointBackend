using Microsoft.AspNetCore.Mvc;

namespace Checkpoint.API.Controllers
{
    [Route("[controller]")]
    [ApiController]
    public class HelloController : ControllerBase
    {
        [HttpGet]
        public IActionResult Hello() {
            return Ok("Hello World!");
        }
    }
}
