unit Unit1;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, StdCtrls, ComCtrls,
  Menus, MaskEdit, CheckLst, Grids, ExtCtrls, uRichEdit;

type

  { TForm1 }

  TForm1 = class(TForm)
    Button1: TButton;
    Button2: TButton;
    Button3: TButton;
    INSzhuangtai: TStringGrid;
    Tiezi_huashu: TEdit;
    PageControl1: TPageControl;
    TabSheet1: TTabSheet;
    TabSheet2: TTabSheet;
    TabSheet3: TTabSheet;
    ZhuangTai: TTimer;
    procedure Button1Click(Sender: TObject);
    procedure Button2Click(Sender: TObject);
    procedure Button3Click(Sender: TObject);
    procedure HostDailiChange(Sender: TObject);
    procedure Timer1Timer(Sender: TObject);
    procedure ZhuangTaiTimer(Sender: TObject);
  private

  public

  end;

var
  Form1: TForm1;

implementation

{$R *.lfm}

{ TForm1 }

procedure TForm1.Button1Click(Sender: TObject);
begin

end;

procedure TForm1.Button2Click(Sender: TObject);
begin

end;

procedure TForm1.Button3Click(Sender: TObject);
begin

end;

procedure TForm1.HostDailiChange(Sender: TObject);
begin

end;

procedure TForm1.Timer1Timer(Sender: TObject);
begin

end;

procedure TForm1.ZhuangTaiTimer(Sender: TObject);
begin

end;

end.

