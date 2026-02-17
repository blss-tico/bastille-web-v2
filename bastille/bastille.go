package bastille

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func RunBastilleCommands(args ...string) (string, error) {
	cmd := exec.Command("bastille", args...)
	log.Println("RunBastilleCommands:", cmd)

	result, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("bastille: %s ,failed: %v\n %s", cmd, err, result)
	}

	return string(result), nil
}

func RunBastilleCommandAsync(args ...string) (<-chan string, <-chan error) {
	resultChan := make(chan string, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(resultChan)
		defer close(errChan)

		cmd := exec.Command("bastille", args...)
		log.Println("runBastilleCommandAsync:", cmd)

		result, err := cmd.CombinedOutput()
		if err != nil {
			errChan <- fmt.Errorf("bastille: %s, failed: %v\n%s", cmd, err, string(result))
			return
		}

		resultChan <- string(result)
	}()

	return resultChan, errChan
}

type Bastille struct{}

func NewBastille() *Bastille {
	return &Bastille{}
}

func (b *Bastille) BastilleVersion() (string, error) {
	args := []string{"-v"}
	return RunBastilleCommands(args...)
}

func (b *Bastille) Bootstrap(options, releasetemplate, updatearch string) (string, error) {
	args := []string{"bootstrap"}
	if options != "" {
		args = append(args, options)
	}

	if releasetemplate != "" {
		args = append(args, releasetemplate)
	}

	if updatearch != "" {
		args = append(args, updatearch)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Clone(options, target, newname, ip string) (string, error) {
	args := []string{"clone"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if newname != "" {
		args = append(args, newname)
	}

	if ip != "" {
		args = append(args, ip)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Cmd(options, target string, command []string) (string, error) {
	args := []string{"cmd"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	args = append(args, command...)

	res, err := RunBastilleCommands(args...)
	return res, err
}

func (b *Bastille) Config(options, target, action, property, value string) (string, error) {
	args := []string{"config"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if action != "" {
		args = append(args, action)
	}

	if property != "" {
		args = append(args, property)
	}

	if value != "" {
		args = append(args, value)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Console(options, target, user string) (string, error) {
	port := RandPortUtil()
	args := []string{"-t", "disableLeaveAlert=true", "-o", "-p", port, "-W", "sudo", "bastille", "console"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if user != "" {
		args = append(args, user)
	}

	cmd := exec.Command("ttyd", args...)
	if err := cmd.Start(); err != nil {
		log.Println("Error starting ttyd:", err)
		return "", fmt.Errorf("bastille: %s ,failed: %v\n", cmd, err)
	}

	return port, nil
}

func (b *Bastille) Convert(options, target, release string) (string, error) {
	args := []string{"convert"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if release != "" {
		args = append(args, release)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Cp(options, target, hostpath, jailpath string) (string, error) {
	args := []string{"cp"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if hostpath != "" {
		args = append(args, hostpath)
	}

	if jailpath != "" {
		args = append(args, jailpath)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Create(options, name, release, ip, iface, gtwip, ipip, value, vlanid, zfsoptions string) (string, error) {
	args := []string{"create"}

	if options != "" {
		args = append(args, options)
	}

	if gtwip != "" {
		gtwipArgs := []string{"-g"}
		gtwipArgs = append(gtwipArgs, gtwip)
		args = append(args, gtwipArgs...)
	}

	if ipip != "" {
		ipipArgs := []string{"-n"}
		ipipArgs = append(ipipArgs, ipip)
		args = append(args, ipipArgs...)
	}

	if value != "" {
		valueArgs := []string{"-p"}
		valueArgs = append(valueArgs, value)
		args = append(args, valueArgs...)
	}

	if vlanid != "" {
		vlanidArgs := []string{"-v"}
		vlanidArgs = append(vlanidArgs, vlanid)
		args = append(args, vlanidArgs...)
	}

	if zfsoptions != "" {
		zfsoptionsArgs := []string{"-Z"}
		zfsoptionsArgs = append(zfsoptionsArgs, zfsoptions)
		args = append(args, zfsoptionsArgs...)
	}

	if name != "" {
		args = append(args, name)
	}

	if release != "" {
		args = append(args, release)
	}

	if ip != "" {
		args = append(args, ip)
	}

	if iface != "" {
		args = append(args, iface)
	}

	log.Println(args)
	return RunBastilleCommands(args...)
}

func (b *Bastille) Destroy(options, jailrelease string) (string, error) {
	args := []string{"destroy"}

	if options != "" {
		args = append(args, options)
	}

	if jailrelease != "" {
		args = append(args, jailrelease)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Edit(options, target, file string) (string, error) {
	port := RandPortUtil()
	args := []string{"-t", "disableLeaveAlert=true", "-o", "-p", port, "-W", "sudo", "bastille", "edit"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if file != "" {
		args = append(args, file)
	}

	cmd := exec.Command("ttyd", args...)
	if err := cmd.Start(); err != nil {
		log.Println("Error starting ttyd:", err)
		return "", fmt.Errorf("bastille: %s ,failed: %v\n", cmd, err)
	}

	return port, nil
}

func (b *Bastille) Etcupdate(options, bootstraptarget, action, release string) (string, error) {
	args := []string{"etcupdate"}

	if options != "" {
		args = append(args, options)
	}

	if bootstraptarget != "" {
		args = append(args, bootstraptarget)
	}

	if action != "" {
		args = append(args, action)
	}

	if release != "" {
		args = append(args, release)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Export(options []string, target, path string) (string, error) {
	args := []string{"export"}

	args = append(args, options...)

	if target != "" {
		args = append(args, target)
	}

	if path != "" {
		args = append(args, path)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Htop(options, target string) (string, error) {
	port := RandPortUtil()
	args := []string{"-t", "disableLeaveAlert=true", "-o", "-p", port, "-W", "sudo", "bastille", "htop"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	cmd := exec.Command("ttyd", args...)
	if err := cmd.Start(); err != nil {
		log.Println("Error starting ttyd:", err)
		return "", fmt.Errorf("bastille: %s ,failed: %v\n", cmd, err)
	}

	return port, nil
}

func (b *Bastille) ImporT(options, file, release string) (string, error) {
	args := []string{"import"}

	if options != "" {
		args = append(args, options)
	}

	if file != "" {
		args = append(args, file)
	}

	if release != "" {
		args = append(args, release)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Jcp(options, sourcejail, jailpath, destjail, jailpath2 string) (string, error) {
	args := []string{"jcp"}

	if options != "" {
		args = append(args, options)
	}

	if sourcejail != "" {
		args = append(args, sourcejail)
	}

	if jailpath != "" {
		args = append(args, jailpath)
	}

	if destjail != "" {
		args = append(args, destjail)
	}

	if jailpath2 != "" {
		args = append(args, jailpath2)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Limits(options, target, action, option, value string) (string, error) {
	args := []string{"limits"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if action != "" {
		args = append(args, action)
	}

	if option != "" {
		args = append(args, option)
	}

	if value != "" {
		args = append(args, value)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) List(options, action string) (string, error) {
	args := []string{"list"}

	if options != "" {
		args = append(args, options)
	}

	if action != "" {
		args = append(args, action)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Migrate(options, target, remote string) (string, error) {
	args := []string{"migrate"}
	passOpt := false

	if options != "" {
		args = append(args, options)

		if strings.Contains(options, "-p") {
			passOpt = true
		}
	}

	if target != "" {
		args = append(args, target)
	}

	if remote != "" {
		args = append(args, remote)
	}

	if passOpt {
		port := RandPortUtil()
		cmd := exec.Command("ttyd", "-t", "disableLeaveAlert=true", "-o", "-p", port, "-W", "sudo", "bastille", "migrate", "-p", target, remote)
		log.Println("console: ", cmd)

		if err := cmd.Start(); err != nil {
			log.Println("Error starting ttyd:", err)
			return "", fmt.Errorf("bastille: %s ,failed: %v\n", cmd, err)
		}

		return port, nil
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Monitor(options, target, action, service string) (string, error) {
	args := []string{"monitor"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if action != "" {
		args = append(args, action)
	}

	if service != "" {
		args = append(args, service)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Mount(options, target, hostpath, jailpath, filesystemtype, option, dump, passnumber string) (string, error) {
	args := []string{"mount"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if hostpath != "" {
		args = append(args, hostpath)
	}

	if jailpath != "" {
		args = append(args, jailpath)
	}

	if filesystemtype != "" {
		args = append(args, filesystemtype)
	}

	if option != "" {
		args = append(args, option)
	}

	if dump != "" {
		args = append(args, dump)
	}

	if passnumber != "" {
		args = append(args, passnumber)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Network(options, target, action, iface, ip, vlanid string) (string, error) {
	args := []string{"network"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if action != "" {
		args = append(args, action)
	}

	if iface != "" {
		args = append(args, iface)
	}

	if ip != "" {
		args = append(args, ip)
	}

	if vlanid != "" {
		vlanidArgs := []string{"-v"}
		vlanidArgs = append(vlanidArgs, vlanid)
		args = append(args, vlanidArgs...)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Pkg(options, target string, arg []string) (string, error) {
	args := []string{"pkg"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	args = append(args, arg...)

	return RunBastilleCommands(args...)
}

func (b *Bastille) Rcp(options, target, jailpath, hostpath string) (string, error) {
	args := []string{"rcp"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if jailpath != "" {
		args = append(args, jailpath)
	}

	if hostpath != "" {
		args = append(args, hostpath)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Rdr(options, odestination, ointerface, osource, otype, target, action, hostport, jailport, logg, logopts string) (string, error) {
	args := []string{"rdr"}

	if options != "" {
		args = append(args, options)
	}

	if odestination != "" {
		odestinationArgs := []string{"-d"}
		odestinationArgs = append(odestinationArgs, odestination)
		args = append(args, odestinationArgs...)
	}

	if ointerface != "" {
		ointerfaceArgs := []string{"-i"}
		ointerfaceArgs = append(ointerfaceArgs, ointerface)
		args = append(args, ointerfaceArgs...)
	}

	if osource != "" {
		osourceArgs := []string{"-s"}
		osourceArgs = append(osourceArgs, osource)
		args = append(args, osourceArgs...)
	}

	if otype != "" {
		otypeArgs := []string{"-t"}
		otypeArgs = append(otypeArgs, otype)
		args = append(args, otypeArgs...)
	}

	if target != "" {
		args = append(args, target)
	}

	if action != "" {
		args = append(args, action)
	}

	if hostport != "" {
		args = append(args, hostport)
	}

	if jailport != "" {
		args = append(args, jailport)
	}

	if logg != "" {
		args = append(args, logg)
	}

	if logopts != "" {
		args = append(args, logopts)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Rename(options, target, newname string) (string, error) {
	args := []string{"rename"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if newname != "" {
		args = append(args, newname)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Restart(options, target, value string) (string, error) {
	args := []string{"restart"}
	if options != "" {
		args = append(args, options)
	}

	if value != "" {
		args = append(args, value)
	}

	if target != "" {
		args = append(args, target)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Service(options, target, servicename, arg string) (string, error) {
	args := []string{"service"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if servicename != "" {
		args = append(args, servicename)
	}

	if arg != "" {
		args = append(args, arg)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Setup(options, action string) (string, error) {
	args := []string{"setup"}
	if options != "" {
		args = append(args, options)
	}

	if action != "" {
		args = append(args, action)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Start(options, target, value string) (string, error) {
	args := []string{"start"}
	if options != "" {
		args = append(args, options)
	}

	if value != "" {
		args = append(args, value)
	}

	if target != "" {
		args = append(args, target)
	}

	//return RunBastilleCommands(args...)
	resultChan, errChan := RunBastilleCommandAsync(args...)
	select {
	case res := <-resultChan:
		return res, nil
	case err := <-errChan:
		return "", err
	}
}

func (b *Bastille) Stop(options, target string) (string, error) {
	args := []string{"stop"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	// return RunBastilleCommands(args...)
	resultChan, errChan := RunBastilleCommandAsync(args...)
	select {
	case res := <-resultChan:
		return res, nil
	case err := <-errChan:
		return "", err
	}
}

func (b *Bastille) Sysrc(options, target, arg string) (string, error) {
	args := []string{"sysrc"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if arg != "" {
		args = append(args, arg)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Tags(options, target, action, tags string) (string, error) {
	args := []string{"tags"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if action != "" {
		args = append(args, action)
	}

	if tags != "" {
		args = append(args, tags)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Template(options, target, action, template string) (string, error) {
	args := []string{"template"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if action != "" {
		args = append(args, action)
	}

	if template != "" {
		args = append(args, template)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Top(options, target string) (string, error) {
	port := RandPortUtil()
	args := []string{"-t", "disableLeaveAlert=true", "-o", "-p", port, "-W", "sudo", "bastille", "top"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	cmd := exec.Command("ttyd", args...)
	if err := cmd.Start(); err != nil {
		log.Println("Error starting ttyd:", err)
		return "", fmt.Errorf("bastille: %s ,failed: %v\n", cmd, err)
	}

	return port, nil
}

func (b *Bastille) Umount(options, target, jailpath string) (string, error) {
	args := []string{"umount"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if jailpath != "" {
		args = append(args, jailpath)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Update(options, target string) (string, error) {
	args := []string{"update"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Upgrade(options, target, action string) (string, error) {
	args := []string{"upgrade"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if action != "" {
		args = append(args, action)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Verify(options, action string) (string, error) {
	args := []string{"verify"}
	if options != "" {
		args = append(args, options)
	}

	if action != "" {
		args = append(args, action)
	}

	return RunBastilleCommands(args...)
}

func (b *Bastille) Zfs(options, target, action, tag, key, value, pooldataset, jailpath string) (string, error) {
	args := []string{"zfs"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if action != "" {
		args = append(args, action)
	}

	if tag != "" {
		args = append(args, tag)
	}

	if key != "" {
		args = append(args, key)
	}

	if value != "" {
		args = append(args, value)
	}

	if pooldataset != "" {
		args = append(args, pooldataset)
	}

	if jailpath != "" {
		args = append(args, jailpath)
	}

	return RunBastilleCommands(args...)
}
