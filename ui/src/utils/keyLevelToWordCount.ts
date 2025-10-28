export default function keyLevelToWordCount(level: number): 6 | 12 | 18 | 24{
    switch(level) {
        case 2 : return 6;
        case 3 : return 12;
        case 4 : return 18;
        case 5 : return 24;
    }
    return 6;
}