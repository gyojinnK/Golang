namespace week05_01
{
    partial class Form1
    {
        /// <summary>
        ///  Required designer variable.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        ///  Clean up any resources being used.
        /// </summary>
        /// <param name="disposing">true if managed resources should be disposed; otherwise, false.</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows Form Designer generated code

        /// <summary>
        ///  Required method for Designer support - do not modify
        ///  the contents of this method with the code editor.
        /// </summary>
        private void InitializeComponent()
        {
            this.panel1 = new System.Windows.Forms.Panel();
            this.label1 = new System.Windows.Forms.Label();
            this.panel2 = new System.Windows.Forms.Panel();
            this.pbTryProgressBar = new System.Windows.Forms.ProgressBar();
            this.lblState = new System.Windows.Forms.Label();
            this.lblTryCount = new System.Windows.Forms.Label();
            this.lblStartTime = new System.Windows.Forms.Label();
            this.label4 = new System.Windows.Forms.Label();
            this.label3 = new System.Windows.Forms.Label();
            this.label2 = new System.Windows.Forms.Label();
            this.panel3 = new System.Windows.Forms.Panel();
            this.btnInit = new System.Windows.Forms.Button();
            this.btnInput = new System.Windows.Forms.Button();
            this.txtInput = new System.Windows.Forms.TextBox();
            this.panel1.SuspendLayout();
            this.panel2.SuspendLayout();
            this.panel3.SuspendLayout();
            this.SuspendLayout();
            // 
            // panel1
            // 
            this.panel1.Controls.Add(this.label1);
            this.panel1.Dock = System.Windows.Forms.DockStyle.Top;
            this.panel1.Location = new System.Drawing.Point(0, 0);
            this.panel1.Name = "panel1";
            this.panel1.Size = new System.Drawing.Size(800, 100);
            this.panel1.TabIndex = 0;
            // 
            // label1
            // 
            this.label1.AutoSize = true;
            this.label1.Location = new System.Drawing.Point(48, 39);
            this.label1.Name = "label1";
            this.label1.Size = new System.Drawing.Size(99, 15);
            this.label1.TabIndex = 0;
            this.label1.Text = "숫자 맞추기 게임";
            // 
            // panel2
            // 
            this.panel2.Controls.Add(this.pbTryProgressBar);
            this.panel2.Controls.Add(this.lblState);
            this.panel2.Controls.Add(this.lblTryCount);
            this.panel2.Controls.Add(this.lblStartTime);
            this.panel2.Controls.Add(this.label4);
            this.panel2.Controls.Add(this.label3);
            this.panel2.Controls.Add(this.label2);
            this.panel2.Dock = System.Windows.Forms.DockStyle.Fill;
            this.panel2.Location = new System.Drawing.Point(0, 100);
            this.panel2.Name = "panel2";
            this.panel2.Size = new System.Drawing.Size(800, 350);
            this.panel2.TabIndex = 1;
            // 
            // pbTryProgressBar
            // 
            this.pbTryProgressBar.Location = new System.Drawing.Point(258, 98);
            this.pbTryProgressBar.Name = "pbTryProgressBar";
            this.pbTryProgressBar.Size = new System.Drawing.Size(502, 23);
            this.pbTryProgressBar.TabIndex = 6;
            // 
            // lblState
            // 
            this.lblState.AutoSize = true;
            this.lblState.Location = new System.Drawing.Point(183, 142);
            this.lblState.Name = "lblState";
            this.lblState.Size = new System.Drawing.Size(191, 15);
            this.lblState.TabIndex = 5;
            this.lblState.Text = "1에서 100사이의 수를 추측하세요";
            // 
            // lblTryCount
            // 
            this.lblTryCount.AutoSize = true;
            this.lblTryCount.Location = new System.Drawing.Point(183, 106);
            this.lblTryCount.Name = "lblTryCount";
            this.lblTryCount.Size = new System.Drawing.Size(0, 15);
            this.lblTryCount.TabIndex = 4;
            // 
            // lblStartTime
            // 
            this.lblStartTime.AutoSize = true;
            this.lblStartTime.Location = new System.Drawing.Point(183, 71);
            this.lblStartTime.Name = "lblStartTime";
            this.lblStartTime.Size = new System.Drawing.Size(0, 15);
            this.lblStartTime.TabIndex = 3;
            // 
            // label4
            // 
            this.label4.AutoSize = true;
            this.label4.Location = new System.Drawing.Point(48, 142);
            this.label4.Name = "label4";
            this.label4.Size = new System.Drawing.Size(62, 15);
            this.label4.TabIndex = 2;
            this.label4.Text = "현재상태 :";
            // 
            // label3
            // 
            this.label3.AutoSize = true;
            this.label3.Location = new System.Drawing.Point(48, 106);
            this.label3.Name = "label3";
            this.label3.Size = new System.Drawing.Size(62, 15);
            this.label3.TabIndex = 1;
            this.label3.Text = "시도횟수 :";
            // 
            // label2
            // 
            this.label2.AutoSize = true;
            this.label2.Location = new System.Drawing.Point(48, 71);
            this.label2.Name = "label2";
            this.label2.Size = new System.Drawing.Size(62, 15);
            this.label2.TabIndex = 0;
            this.label2.Text = "시작시간 :";
            // 
            // panel3
            // 
            this.panel3.Controls.Add(this.btnInit);
            this.panel3.Controls.Add(this.btnInput);
            this.panel3.Controls.Add(this.txtInput);
            this.panel3.Dock = System.Windows.Forms.DockStyle.Bottom;
            this.panel3.Location = new System.Drawing.Point(0, 350);
            this.panel3.Name = "panel3";
            this.panel3.Size = new System.Drawing.Size(800, 100);
            this.panel3.TabIndex = 2;
            // 
            // btnInit
            // 
            this.btnInit.Location = new System.Drawing.Point(48, 33);
            this.btnInit.Name = "btnInit";
            this.btnInit.Size = new System.Drawing.Size(75, 23);
            this.btnInit.TabIndex = 2;
            this.btnInit.Text = "게임시작";
            this.btnInit.UseVisualStyleBackColor = true;
            this.btnInit.Click += new System.EventHandler(this.btnInit_Click);
            // 
            // btnInput
            // 
            this.btnInput.Location = new System.Drawing.Point(481, 33);
            this.btnInput.Name = "btnInput";
            this.btnInput.Size = new System.Drawing.Size(75, 23);
            this.btnInput.TabIndex = 1;
            this.btnInput.Text = "입력";
            this.btnInput.UseVisualStyleBackColor = true;
            this.btnInput.Click += new System.EventHandler(this.btnInput_Click);
            // 
            // txtInput
            // 
            this.txtInput.Location = new System.Drawing.Point(346, 33);
            this.txtInput.Name = "txtInput";
            this.txtInput.Size = new System.Drawing.Size(100, 23);
            this.txtInput.TabIndex = 0;
            this.txtInput.KeyDown += new System.Windows.Forms.KeyEventHandler(this.txtInput_KeyDown);
            // 
            // Form1
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(7F, 15F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(800, 450);
            this.Controls.Add(this.panel3);
            this.Controls.Add(this.panel2);
            this.Controls.Add(this.panel1);
            this.Name = "Form1";
            this.Text = "5주차 오픈소스프로그래밍";
            this.panel1.ResumeLayout(false);
            this.panel1.PerformLayout();
            this.panel2.ResumeLayout(false);
            this.panel2.PerformLayout();
            this.panel3.ResumeLayout(false);
            this.panel3.PerformLayout();
            this.ResumeLayout(false);

        }

        #endregion

        private Panel panel1;
        private Label label1;
        private Panel panel2;
        private Label lblState;
        private Label lblTryCount;
        private Label lblStartTime;
        private Label label4;
        private Label label3;
        private Label label2;
        private Panel panel3;
        private Button btnInit;
        private Button btnInput;
        private TextBox txtInput;
        private ProgressBar pbTryProgressBar;
    }
}