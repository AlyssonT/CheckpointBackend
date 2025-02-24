namespace Checkpoint.Exceptions.ExceptionBase;

public class NotFoundException : CheckpointException
{
    public override string Message { get; }
    public NotFoundException(string message)
    {
        Message = message;
    }
}
