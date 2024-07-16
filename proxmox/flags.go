package proxmox

import (
	"context"

	"github.com/rancher/machine/libmachine/drivers"
	"github.com/rancher/machine/libmachine/mcnflag"
)

var qemuFlags = []mcnflag.Flag{
	mcnflag.StringFlag{
		EnvVar: "PROXMOX_QEMU_METHOD",
		Name:   "proxmox-qemu-method",
		Usage:  "Method to put the ssh credentials from the box: agent (qemu-guest-agent), drive (PVE cloud-init drive), nocloud (cloud-init iso)",
	},
	mcnflag.StringFlag{
		EnvVar: "PROXMOX_QEMU_URL",
		Name:   "proxmox-qemu-url",
		Usage:  "URL to the proxmox API Server",
	},
	mcnflag.StringFlag{
		EnvVar: "PROXMOX_QEMU_USERNAME",
		Name:   "proxmox-qemu-username",
		Usage:  "Username for the proxmox API Server",
	},
	mcnflag.StringFlag{
		EnvVar: "PROXMOX_QEMU_PASSWORD",
		Name:   "proxmox-qemu-password",
		Usage:  "Password for the proxmox API Server",
	},
	mcnflag.StringFlag{
		EnvVar: "PROXMOX_QEMU_TOKENID",
		Name:   "proxmox-qemu-tokenid",
		Usage:  "Token ID for the proxmox API Server",
	},
	mcnflag.StringFlag{
		EnvVar: "PROXMOX_QEMU_SECRET",
		Name:   "proxmox-qemu-secret",
		Usage:  "Secret for a TokenID for the proxmox API Server",
	},
	mcnflag.StringFlag{
		EnvVar: "PROXMOX_QEMU_2FA_CODE",
		Name:   "proxmox-qemu-2fa-code",
		Usage:  "Two Factor Authentication code for logins, not required if 2fa not turned on",
	},
	mcnflag.BoolFlag{
		EnvVar: "PROXMOX_QEMU_INSECURE",
		Name:   "proxmox-qemu-insecure",
		Usage:  "Skip TLS verification",
	},
	mcnflag.IntFlag{
		EnvVar: "PROXMOX_QEMU_TIMEOUT",
		Name:   "proxmox-qemu-timeout",
		Usage:  "API timeout in seconds",
		Value:  30,
	},
	mcnflag.StringFlag{
		EnvVar: "PROXMOX_QEMU_NODE",
		Name:   "proxmox-qemu-node",
		Usage:  "Node name the template is on",
	},
	mcnflag.IntFlag{
		EnvVar: "PROXMOX_QEMU_TEMPLATE_ID",
		Name:   "proxmox-qemu-template-id",
		Usage:  "Id of the template to clone from",
	},
}

func (d *QemuDriver) SetConfigFromFlags(opts drivers.DriverOptions) error {
	d.Method = opts.String("proxmox-qemu-method")
	d.ApiUrl = opts.String("proxmox-qemu-url")
	d.Username = opts.String("proxmox-qemu-username")
	d.Password = opts.String("proxmox-qemu-password")
	d.TwoFactorAuthCode = opts.String("proxmox-qemu-2fa-code")
	d.Insecure = opts.Bool("proxmox-qemu-insecure")
	d.TemplateId = opts.Int("proxmox-qemu-template-id")
	d.Node = opts.String("proxmox-qemu-node")
	d.TokenID = opts.String("proxmox-qemu-tokenid")
	d.Secret = opts.String("proxmox-qemu-secret")
	d.client = d.proxmoxClient()

	ctx := context.Background()
	_, err := d.client.Version(ctx) // get version info to verify credentials

	return err
}
