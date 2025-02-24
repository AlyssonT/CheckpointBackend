using System;

namespace Checkpoint.Exceptions.ExceptionBase;

public class UserAlreadyExistsException : CheckpointException
{
    public UserAlreadyExistsException(string message) : base(message)
    {
    }
}