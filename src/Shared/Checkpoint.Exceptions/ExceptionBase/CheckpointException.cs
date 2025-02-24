namespace Checkpoint.Exceptions.ExceptionBase;

public class CheckpointException : SystemException
{
    public override string Message { get; }
    public CheckpointException(string message)
    {
        Message = message;
    }
}
