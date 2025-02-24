namespace Checkpoint.Exceptions.ExceptionBase;

public class NotFoundException : CheckpointException
{
    public NotFoundException(string message) : base(message)
    {
    }
}
