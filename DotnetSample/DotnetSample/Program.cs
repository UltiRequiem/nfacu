using System;
using System.Collections.Generic;
using System.Configuration;
using System.Linq;

namespace DotnetSample
{
    internal class Program
    {
        private static void Main()
        {
            foreach (var (item1, item2) in GetConfigurationValues())
            {
                Console.WriteLine($"{item1} => {item2}");
            }

            Console.ReadKey();
        }

        public static IEnumerable<Tuple<string, string>> GetConfigurationValues() => ConfigurationManager.AppSettings
           .AllKeys.Select(key => Tuple.Create(key, ConfigurationManager.AppSettings[key])).ToList();
    }
}
