#include <stdio.h>
#include <stdlib.h>
#include <conio.h>
#include <Windows.h>

#define LINES 15

HANDLE console;
CHAR_INFO *buffer;
COORD bufferSize;

void init()
{
    system("chcp 65001");

    console = GetStdHandle(STD_OUTPUT_HANDLE);
    CONSOLE_SCREEN_BUFFER_INFO consoleInfo;
    GetConsoleScreenBufferInfo(console, &consoleInfo);
    bufferSize = consoleInfo.dwSize;

    buffer = (CHAR_INFO *)malloc(bufferSize.X * bufferSize.Y * sizeof(CHAR_INFO));
}

void clear()
{
}

void draw()
{
    memset(buffer, 0, bufferSize.X * bufferSize.Y * sizeof(CHAR_INFO));
    char *map[LINES][LINES + 1];
    for (int i = 0; i < LINES; i++)
    {
        int j;
        for (j = 0; j < LINES; j++)
        {
            if (i == 0)
            {
                if (j == 0)
                {
                    // strcpy(map[i][j], "┌");
                    buffer[i * (LINES + 1) + j].Char.AsciiChar = '┌';
                }
                else if (j == LINES - 1)
                {
                    // strcpy(map[i][j], '┐');
                    buffer[i * (LINES + 1) + j].Char.AsciiChar = '┐';
                }
                else
                {
                    // strcpy(map[i][j], '┬');
                    buffer[i * (LINES + 1) + j].Char.AsciiChar = '┬';
                }
            }
            else if (i == LINES - 1)
            {
                if (j == 0)
                {
                    // strcpy(map[i][j], '└');
                    buffer[i * (LINES + 1) + j].Char.AsciiChar = '└';
                }
                else if (j == LINES - 1)
                {
                    // strcpy(map[i][j], '┘');
                    buffer[i * (LINES + 1) + j].Char.AsciiChar = '┘';
                }
                else
                {
                    // strcpy(map[i][j], '┴');
                    buffer[i * (LINES + 1) + j].Char.AsciiChar = '┴';
                }
            }
            else
            {
                if (j == 0)
                {
                    // strcpy(map[i][j], '├');
                    buffer[i * (LINES + 1) + j].Char.AsciiChar = '├';
                }
                else if (j == LINES - 1)
                {
                    // strcpy(map[i][j], '┤');
                    buffer[i * (LINES + 1) + j].Char.AsciiChar = '┤';
                }
                else
                {
                    // strcpy(map[i][j], '┼');
                    buffer[i * (LINES + 1) + j].Char.AsciiChar = '┼';
                }
            }
        }
        // strcpy(map[i][LINES], '\n');
        buffer[i * (LINES + 1) + j].Char.AsciiChar = '\n';
    }

    // for (int i = 0; i < LINES; i++)
    // {
    //     for (int j = 0; j <= LINES; j++)
    //     {
    //         buffer[i * (LINES + 1) + j].Char.UnicodeChar = (WCHAR)map[i][j];
    //     }
    // }

    COORD bufferCoord = {0, 0};
    SMALL_RECT writeRegion = {0, 0, bufferSize.X - 1, bufferSize.Y - 1};
    WriteConsoleOutput(console, buffer, bufferSize, bufferCoord, &writeRegion);
    FlushConsoleInputBuffer(console);
}

void processInput()
{
}

void update() {}

int main(void)
{
    init();

    while (1)
    {
        processInput();
        update();
        draw();
    }

    clear();

    return 0;
}