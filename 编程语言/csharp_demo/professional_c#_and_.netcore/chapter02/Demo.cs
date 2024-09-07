
#define DEBUG
#undef RELEASE


namespace Nico.Chapter02
{
    ///<summary>
    /// Xml document
    ///</summary>
    public class Demo
    {
        public static int Display()
        {
#if DEBUG
            Console.WriteLine("Debug mode!");
#endif

            return 0;
        }

#pragma warning disable
        private static void Type()
        {
            #region inttype

            sbyte sb;
            short s;
            int i;
            long l;
            byte b;
            ushort us;
            uint ui;
            ulong ul;

            #endregion


            float f;
            double d;
            decimal dc;

            bool bo;

            char c;

            object o;

            string str = "Hello";
            Console.WriteLine($"str is {str}");

            const int j = 0;
        }

        private static void If(int i)
        {
            if (i > 0)
            {
                Console.WriteLine("bigger than zero");
            }
            else if (i < 0)
            {
                Console.WriteLine("less than zero");
            }
            else
            {
                Console.WriteLine("equals to zero");
            }
        }

        private static void Switch(string s)
        {
            switch (s)
            {
                case "nico":
                    Console.WriteLine("Hello nico");
                    break;
                case "a":
                case "b":
                    Console.WriteLine("s = " + s);
                    break;
                default:
                    Console.WriteLine("others " + s);
                    break;
            }
        }

        private static void For()
        {
            for (int i = 0; i < 10; i++)
            {
                Console.WriteLine(i);
            }

            int[] arr = [1, 3, 5];
            foreach (var i in arr)
            {
                Console.WriteLine(i);
            }
        }

        private static void While()
        {
            int i = 10;
            while (i > 0)
            {
                Console.WriteLine(i--);
            }

            i = 10;
            do
            {
                Console.WriteLine(i--);
            } while (i > 0);
        }
    }

}