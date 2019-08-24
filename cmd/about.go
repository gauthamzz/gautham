/*
Copyright © 2019 GAUTHAM SANTHOSH thabeatsz@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// aboutCmd represents the about command
var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("👋  Hey, nice to meet you! \n I'm Gautham Santhosh, a programmer, designer and Nomad. \n\n I'm 23 years old. 💛 \n\n I was a student at Indian Institute of Information Technology, Allahabad for integrated M-Tech in Intelligent Systems. \n\n I am originally from Mannanam, Kottayam, Kerala, India. 🇮🇳 \n\nI love building new things and helping others build things. I had Co-Founded one of India's best Student Hackathon, Hack in the north , which i had been organising for the past 3 years. I had also founded a Intra-Allahabad Hackathon, IIIT Hacks .\n\n 🏗 I had previously built Anna Assistant - a virtual assistant for web browsers and Machine Learning Jobs List , a machine learning jobs portal, Pelican Facebook , a facebook experince minimal and focused optimising extension.\n\n At my college, I was Co-ordinator of the Free and Open Source Society, where i had started OpenCode a Open Source event to help people start developement. Where we teach more than 200+ hackers from all over - WebDev, AppDev, Machine Learning and Competitive Coding during Jan-Feb every year from 2017. 👨🎓 I was also Editor in Cheif and Founder of Nybles , a tech news publication. 👽 I love Deep Learning and think it's the future. I had done some work in Generative Models, such as Audio Style Transfer , made Hazel which allows users to build, train, and ship custom deep learning models using a simple visual interface and love creating chatbots such as Indian Lawyer Bot, a Chat bot that helps assists you in Criminal Cases in India. I also advocate solving problems without code, hence sell yo - a platform i had build to buy and sell things between people you trust was made, Trust within automation, hence my new found interest in blockchain. I love watching highly rated boring movies, reading books(non fiction), brainstroming new solutions for problems. My resume is avaliable here . 🔆 Almost everything is opensourced and you can find it on github . \n\n\n You can email me at thabeatsz@gmail.com. ✉️")
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aboutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aboutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
