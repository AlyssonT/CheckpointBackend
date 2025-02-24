using System.Security.Cryptography;
using System.Text;

namespace Checkpoint.Application.Services;

public class PasswordEncrypter
{
    private readonly string _additionalKey;
    public PasswordEncrypter(string additionalKey) {
        _additionalKey = additionalKey;
    }
    public string Encrypt(string password)
    {
        var newPassword = password + _additionalKey;

        var bytes = Encoding.UTF8.GetBytes(newPassword);
        var hashBytes = SHA512.HashData(bytes);

        return ConvertBytesToString(hashBytes);
    }

    private string ConvertBytesToString(byte[] hashBytes)
    {
        var stringBuilder = new StringBuilder();
        foreach (var hashByte in hashBytes)
        {
            stringBuilder.Append(hashByte.ToString("X2"));
        }

        return stringBuilder.ToString();
    }
}
