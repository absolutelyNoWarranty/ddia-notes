import glob
import pathlib
import random

data = []
for chapter in glob.glob("chapter*"):
    questions = []
    p_ques = pathlib.Path(chapter) / "questions.md"
    if p_ques.exists():
        question = None
        options = []
        with p_ques.open('r') as f:
            for line in f:
                if line.startswith("Q:"):
                    if question is not None:
                        questions.append([question, "\n".join(options)])
                        options = []
                    question = line.strip()
                else:
                    options.append(line.strip())

    answers = []

    p_ans = pathlib.Path(chapter) / "answers.md"
    if p_ans.exists():
        
        question = None
        ans = []
        with p_ans.open('r') as f:
            for line in f:
                if line.startswith("Q:"):
                    if question is not None:
                        answers.append([question, "\n".join(ans)])
                        ans = []
                    question = line.strip()
                else:
                    ans.append(line.strip())
    if len(questions) == len(answers):
        print(list(zip(questions, answers)))
        for ((q,options), (q_, a)) in zip(questions, answers):
            data.append((q, options, a))
    else:
        print("NOT same length!")
        print(len(questions))
        print(len(answers))
 
while True:
    choice = random.choice(data)
    print(choice[0])
    print(choice[1])
    user_ans = input("> ")
    print(choice[2])
