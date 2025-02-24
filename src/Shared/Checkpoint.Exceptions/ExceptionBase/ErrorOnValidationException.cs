namespace Checkpoint.Exceptions.ExceptionBase;

public class ErrorOnValidationException : CheckpointException
{
    public IList<string> ErrorMessages { get; set; }
    public ErrorOnValidationException(IList<string> errorMessages)
    {
        ErrorMessages = errorMessages;
    }
}
