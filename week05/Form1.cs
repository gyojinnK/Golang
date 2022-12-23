using System.Runtime.Intrinsics.X86;

namespace week05_01
{
    public partial class Form1 : Form
    {
        int target; // 컴퓨터가 발생시킨 난수가 저장된 변수
        int tryCount; // 시도 횟수 저장용 변수
        public Form1()
        {
            InitializeComponent();
            btnInput.Enabled = false;
        }

        private void btnInput_Click(object sender, EventArgs e)
        {
            int number = int.Parse(txtInput.Text);
            tryCount++;
            IncreaseTryProgressBar(tryCount);
            lblTryCount.Text = tryCount.ToString();

            if (number == target){
                lblState.Text = "정답 " + target + "입니다";
            }
            else if (number > target) {
                lblState.Text = "정답은 입력한 수 보다 작습니다. Down!";
                txtInput.Text = "";
            }
            else
            {
                lblState.Text = "정답은 입력한 수 보다 큽니다. Up!";
                txtInput.Text = "";
            }
            if (tryCount == 10)
                CloseGame();
            txtInput.Focus();
        }

        private void btnInit_Click(object sender, EventArgs e)
        {
            btnInput.Enabled = true;
            lblStartTime.Text = DateTime.Now.ToString();
            lblState.Text = "게임을 시작합니다. 1~100 사이의 수를 추측하세요";
            tryCount = 0;
            lblTryCount.Text = tryCount.ToString() + "회";
            txtInput.Text = "";

            Random r = new Random();
            target = r.Next(1, 100);
        }

        private void IncreaseTryProgressBar(int cnt)
        {
            pbTryProgressBar.Value = cnt * 10;
        }
        private void CloseGame()
        {
            lblState.Text = "게임 종료. 정답은 " + target;
            txtInput.Text = "";
            btnInput.Enabled = false;
        }

        private void txtInput_KeyDown(object sender, KeyEventArgs e)
        {
            if(e.KeyCode == Keys.Enter)
                btnInput_Click(sender, e);
        }
    }
}