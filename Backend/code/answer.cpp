#include <iostream>
#include <fstream>
#include <string>
using namespace std;
class cow
{
int num,ts,te,b;     
public:
cow():num(0),ts(0),te(0),b(0) {}
void set_data(int n,int tts, int tte, int bb)
{
num=n;ts=tts;te=tte;b=bb;
}
void set_cow(int n) {num=n;}
void set_ts(int tts){ts=tts;}
void set_te(int tte){te=tte;}
void set_b(int bb){b=bb;}
int get_cow() {return num;}
int get_ts(){return ts;}
int get_te(){return te;}
int get_b(){return b;}
void show_data()
{
cout<<"Cow # "<<num<<endl;
cout<<"Start milking : "<<ts<<endl;
cout<<"End milking : "<<te<<endl;
cout<<"# of blist : "<<b<<endl;
}
};
int main()
{
int N;
int Ts[101],Te[101],B[101];
ifstream fin("blist.in");
fin>>N;
cow cow1[N];
for (int i=0;i<N;i++)
{
fin>>Ts[i]>>Te[i]>>B[i];
cow1[i].set_cow(i);
cow1[i].set_ts(Ts[i]);
cow1[i].set_te(Te[i]);
cow1[i].set_b(B[i]);
}
//for (int i=1;i<N;i++)
//cow1[i].show_data();
int max_buk = 0;
for(int time=1;time<=1000;time++)
{
int buk_at_this_time=0;
for(int i=0;i<N;i++)
{
if(cow1[i].get_ts() <= time && time <= cow1[i].get_te())
{
buk_at_this_time +=cow1[i].get_b();
}
}
max_buk = max(max_buk,buk_at_this_time);
}
ofstream fout("blist.out");
fout <<max_buk<<endl;
return 0;
}