using System.Runtime.Intrinsics.X86;

namespace week05_01
{
    public partial class Form1 : Form
    {
        int target; // ��ǻ�Ͱ� �߻���Ų ������ ����� ����
        int tryCount; // �õ� Ƚ�� ����� ����
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
                lblState.Text = "���� " + target + "�Դϴ�";
            }
            else if (number > target) {
                lblState.Text = "������ �Է��� �� ���� �۽��ϴ�. Down!";
                txtInput.Text = "";
            }
            else
            {
                lblState.Text = "������ �Է��� �� ���� Ů�ϴ�. Up!";
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
            lblState.Text = "������ �����մϴ�. 1~100 ������ ���� �����ϼ���";
            tryCount = 0;
            lblTryCount.Text = tryCount.ToString() + "ȸ";
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
            lblState.Text = "���� ����. ������ " + target;
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