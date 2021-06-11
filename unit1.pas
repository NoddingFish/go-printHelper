unit Unit1;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, Menus, ExtCtrls,
  StdCtrls, Buttons, DBCtrls;

type

  { TForm1 }

  TForm1 = class(TForm)
    Button1: TButton;
    Button2: TButton;
    LogClear: TButton;
    LogBox: TListBox;
    Nick: TEdit;
    SubNick: TEdit;
    Label1: TLabel;
    Label2: TLabel;
    Panel1: TPanel;
    Panel2: TPanel;
    procedure Button1Click(Sender: TObject);
    procedure DBListBox1Click(Sender: TObject);
    procedure FormCreate(Sender: TObject);
    procedure Label1Click(Sender: TObject);
    procedure logBoxClick(Sender: TObject);
    procedure logClick(Sender: TObject);
  private

  public

  end;

var
  Form1: TForm1;

implementation

{$R *.lfm}

{ TForm1 }

procedure TForm1.FormCreate(Sender: TObject);
begin

end;

procedure TForm1.Button1Click(Sender: TObject);
begin
  //showmessage('你好！'+Nick.text+' '+SubNick.text);
  LogBox.Items.Clear;
  LogBox.Items.Add('恭喜你，登录成功：'+Nick.text+' '+SubNick.text);
  LogBox.Items.Add('First line');
  LogBox.Items.Add('Line with random number '+IntToStr(Random(100)));
  LogBox.Items.Add('Third line');
  LogBox.Items.Add('Even a random number '+IntToStr(Random(100)));
end;

procedure TForm1.DBListBox1Click(Sender: TObject);
begin

end;

procedure TForm1.Label1Click(Sender: TObject);
begin

end;

procedure TForm1.logBoxClick(Sender: TObject);
begin

end;


procedure TForm1.logClick(Sender: TObject);
begin

end;

end.

